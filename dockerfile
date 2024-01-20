FROM jenkins/ssh-agent:alpine-jdk17
# FROM jenkins/agent:alpine-jdk17
USER root
COPY --from=golang:1.13-alpine /usr/local/go/ /usr/local/go/
ENV PATH="/usr/local/go/bin:${PATH}"
USER jenkins
