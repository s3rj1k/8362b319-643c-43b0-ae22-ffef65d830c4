{% set wg = pillar.get('wireguard', {}) %}
{% set server = wg.get('server', {}) %}
{% set peers = server.get('peers', []) %}

install_wireguard:
  pkg.installed:
    - name: wireguard

manage_privatekey:
  file.managed:
    - name: /etc/wireguard/privatekey
    - source: salt://overlay/etc/wireguard/privatekey.jinja
    - user: root
    - group: root
    - mode: 600
    - template: jinja
    - context:
        private_key: {{ server.get('private_key') }}

manage_publickey:
  file.managed:
    - name: /etc/wireguard/publickey
    - source: salt://overlay/etc/wireguard/publickey.jinja
    - user: root
    - group: root
    - mode: 600
    - template: jinja
    - context:
        public_key: {{ server.get('public_key') }}

manage_wg0.conf:
  file.managed:
    - name: /etc/wireguard/wg0.conf
    - source: salt://overlay/etc/wireguard/wg0.conf.jinja
    - user: root
    - group: root
    - mode: 600
    - require:
      - pkg: install_wireguard
    - watch_in:
      - service: restart_wireguard_service
    - template: jinja
    - context:
        address: {{ server.get('address') }}
        listen_port: {{ server.get('listen_port') }}
        private_key: {{ server.get('private_key') }}
        peers: {{ peers | json }}

restart_wireguard_service:
  service.running:
    - name: wg-quick@wg0
    - enable: True
    - watch:
      - file: /etc/wireguard/wg0.conf
