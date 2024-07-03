FROM alpine:3.20.1

COPY build/report-generator-linux-amd64 /report-generator

RUN mkdir /reports

ENV CSV_DBUSER postgres
ENV CSV_DBNAME postgres
ENV CSV_REPORTS_PATH /reports

ENTRYPOINT ["/report-generator"]



