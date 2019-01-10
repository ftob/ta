FROM prom/prometheus

COPY config.yml .

ENTRYPOINT [ "/bin/prometheus", "-config.file=config,yml" ]