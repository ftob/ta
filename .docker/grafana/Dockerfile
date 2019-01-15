FROM grafana/grafana:latest

USER root

COPY ./etc/grafana/datasources /usr/share/grafana/conf/provisioning/datasources
COPY ./etc/grafana/dashboards /usr/share/grafana/conf/provisioning/dashboards

EXPOSE 3000

