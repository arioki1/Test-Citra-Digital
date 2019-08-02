FROM alpine
LABEL maintainer="yogainformatika@gmail.com"
LABEL docker_hub="https://hub.docker.com/r/mryoga12345/test_citra_digital"
RUN mkdir /app
WORKDIR /app
COPY Soal_No_2 /app
CMD ["./Soal_No_2"]
EXPOSE 3000