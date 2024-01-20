{% set address_cidr = pillar['wireguard']['server']['address'] %}
{% set ip_address, network_prefix = address_cidr.split('/') | map('trim') | list %}
{% set proxy_domains = pillar['wireguard']['server']['proxy_domains'] %}

install_pdns_recursor:
  pkg.installed:
    - name: pdns-recursor
    - require:
      - sls: wg

enable_pdns_recursor_service:
  service.running:
    - name: pdns-recursor
    - enable: True
    - require:
      - pkg: install_pdns_recursor

manage_override.lua:
  file.managed:
    - name: /etc/powerdns/override.lua
    - source: salt://overlay/etc/powerdns/override.lua.jinja
    - user: root
    - group: root
    - mode: 644
    - require:
      - pkg: install_pdns_recursor
    - watch_in:
      - service: reload_pdns_recursor_service
    - template: jinja
    - context:
        address: {{ ip_address }}
        proxy_domains: {{ proxy_domains }}

manage_override.conf:
  file.managed:
    - name: /etc/powerdns/recursor.d/override.conf
    - source: salt://overlay/etc/powerdns/recursor.d/override.conf.jinja
    - user: root
    - group: root
    - mode: 644
    - require:
      - pkg: install_pdns_recursor
    - watch_in:
      - service: reload_pdns_recursor_service
    - template: jinja
    - context:
        address: {{ ip_address }}
        network_prefix: {{ network_prefix }}

reload_pdns_recursor_service:
  service.running:
    - name: pdns-recursor
    - enable: true
    - watch:
      - file: /etc/powerdns/override.lua
      - file: /etc/powerdns/recursor.d/override.conf
    - onchanges:
      - file: /etc/powerdns/override.lua
      - file: /etc/powerdns/recursor.d/override.conf
