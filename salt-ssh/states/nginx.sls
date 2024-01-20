{% set address_cidr = pillar['wireguard']['server']['address'] %}
{% set ip_address, network_prefix = address_cidr.split('/') | map('trim') | list %}
{% set proxy_domains = pillar['wireguard']['server']['proxy_domains'] %}

install_nginx:
  pkg.installed:
    - name: nginx
    - require:
      - sls: wg

nginx_service:
  service.running:
    - name: nginx
    - enable: True
    - require:
      - pkg: install_nginx

manage_nginx_conf:
  file.managed:
    - name: /etc/nginx/nginx.conf
    - source: salt://overlay/etc/nginx/nginx.conf.jinja
    - user: root
    - group: root
    - mode: 644
    - require:
      - pkg: install_nginx
    - watch_in:
      - service: nginx_service_reload
    - template: jinja
    - context:
        address: {{ ip_address }}
        proxy_domains: {{ proxy_domains }}

nginx_service_reload:
  service.running:
    - name: nginx
    - enable: true
    - watch:
      - file: manage_nginx_conf
    - reload: True
