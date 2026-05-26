FROM python:bullseye
RUN apt-get update \
    && apt-get -y install libpq-dev gcc \
    && pip install psycopg2
COPY . /data
WORKDIR /data
CMD ["pip", "install", "-r", "requirements.txt"]
ENTRYPOINT ["python3", "src/main.py"]