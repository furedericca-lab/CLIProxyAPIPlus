#!/usr/bin/env python3
"""Check whether a candidate CLIProxyAPIPlus management password is accepted."""

from __future__ import annotations

import argparse
import json
import ssl
import sys
import urllib.error
import urllib.parse
import urllib.request


DEFAULT_PASSWORD = "change-me-to-a-strong-password"


def build_probe_url(base_url: str) -> str:
    base_url = base_url.strip().rstrip("/")
    parsed = urllib.parse.urlparse(base_url)
    if parsed.scheme not in {"http", "https"} or not parsed.netloc:
        raise ValueError("URL must include http:// or https:// and a host")
    return f"{base_url}/v0/management/debug"


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Test a candidate password against CLIProxyAPIPlus management authentication."
    )
    parser.add_argument(
        "--url",
        default="http://127.0.0.1:8317",
        help="CLIProxyAPIPlus base URL (default: %(default)s)",
    )
    parser.add_argument(
        "--password",
        default=None,
        help=(
            "candidate password; omitted or empty uses "
            f"{DEFAULT_PASSWORD!r}"
        ),
    )
    parser.add_argument(
        "--timeout",
        type=float,
        default=5.0,
        help="request timeout in seconds (default: %(default)s)",
    )
    parser.add_argument(
        "--insecure",
        action="store_true",
        help="disable TLS certificate verification for HTTPS",
    )
    return parser.parse_args()


def main() -> int:
    args = parse_args()
    password = args.password.strip() if args.password is not None else ""
    if not password:
        password = DEFAULT_PASSWORD

    try:
        probe_url = build_probe_url(args.url)
    except ValueError as exc:
        print(f"ERROR: {exc}", file=sys.stderr)
        return 2

    request = urllib.request.Request(
        probe_url,
        headers={
            "X-Management-Key": password,
            "Accept": "application/json",
            "User-Agent": "cliproxy-management-password-check/1.0",
        },
    )
    context = ssl._create_unverified_context() if args.insecure else None

    try:
        with urllib.request.urlopen(
            request, timeout=args.timeout, context=context
        ) as response:
            body = response.read(4096)
            if response.status == 200:
                print("SUCCESS: the candidate password was accepted.")
                return 0
            print(f"UNEXPECTED: server returned HTTP {response.status}.")
            if body:
                print(body.decode("utf-8", errors="replace"))
            return 3
    except urllib.error.HTTPError as exc:
        body = exc.read(4096).decode("utf-8", errors="replace")
        try:
            message = json.loads(body).get("error", body)
        except json.JSONDecodeError:
            message = body

        if exc.code == 401:
            print("REJECTED: the candidate password was not accepted.")
            return 1
        if exc.code == 403:
            print(f"BLOCKED: {message or 'remote management is disabled or the IP is banned.'}")
            return 3
        if exc.code == 404:
            print("UNAVAILABLE: the management API is disabled or this is not CLIProxyAPIPlus.")
            return 3
        print(f"ERROR: server returned HTTP {exc.code}: {message}")
        return 3
    except urllib.error.URLError as exc:
        print(f"CONNECTION ERROR: {exc.reason}", file=sys.stderr)
        return 4
    except TimeoutError:
        print("CONNECTION ERROR: request timed out.", file=sys.stderr)
        return 4


if __name__ == "__main__":
    raise SystemExit(main())
