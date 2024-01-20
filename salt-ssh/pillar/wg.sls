wireguard:
  server:
    address: "10.42.0.1/24"
    listen_port: 51820
    private_key: "KJMImX1UlFs2JJi0clL9Lf8ecSCs+MqsOA4gErzp1WU="
    public_key: "GI/4YhKPrUwg5mQwBd1l+ZnXT1c2mzexGJc7S1sYIBk="
    peers:
      - public_key: "02+0GghZUfjPYaDJpikrdbC8BytnEInfo8TpU6Ys1H8="
        allowed_ips: "10.42.0.71/24"
    proxy_domains:
      - netflix.com
      - netflix.net
      - nflxext.com
      - nflximg.com
      - nflximg.net
      - nflxso.net
      - nflxvideo.net
      - ifconfig.io
