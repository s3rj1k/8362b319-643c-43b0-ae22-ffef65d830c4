# Wireguard VPN Client (Peer)

## Setup

- Put `./overlay/etc/wireguard/wg0.conf` to your systemd based Linux system into `/etc/wireguard/wg0.conf`.
- Install `wireguard` or `wireguard-tools` and `systemd-resolvconf`.
- Run `sudo systemctl enable wg-quick@wg0.service`.
- Run `sudo systemctl restart wg-quick@wg0.service`.
- Run `sudo wg` to check tunnel status.
- Run `curl --interface wg0 https://ifconfig.io` or `curl https://ifconfig.io` (depends on local DNS server setup) and `curl ifconfig.me` to check public IP difference.
