manage_ipv4-forward.conf:
  file.managed:
    - name: /etc/sysctl.d/ipv4-forward.conf
    - source: salt://overlay/etc/sysctl.d/ipv4-forward.conf
    - user: root
    - group: root
    - mode: '0644'
    - require_in:
      - cmd: reload_sysctl

reload_sysctl:
  cmd.wait:
    - name: sysctl --system
    - watch:
      - file: /etc/sysctl.d/ipv4-forward.conf
