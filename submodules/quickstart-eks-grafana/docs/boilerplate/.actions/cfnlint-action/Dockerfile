FROM python:3.12-buster
RUN pip install cfn-lint
COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
