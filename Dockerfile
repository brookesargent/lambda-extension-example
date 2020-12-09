FROM amazon/aws-lambda-go

ARG HONEYCOMB_WRITE_KEY=""
ARG FUNCTION_DIR="/var/task"
ARG EXTENSION_DIR="/opt/extensions"

# Honeycomb extension
RUN mkdir -p /opt/extensions
COPY honeycomb-lambda-extension ${EXTENSION_DIR}
RUN chmod +x ${EXTENSION_DIR}/honeycomb-lambda-extension

# function code
RUN mkdir -p ${FUNCTION_DIR}
COPY helloWorld ${FUNCTION_DIR}

# set env vars
ENV LIBHONEY_API_KEY HONEYCOMB_WRITE_KEY
ENV LIBHONEY_DATASET lambda_extension_test

# Set the CMD to your handler (could also be done as a parameter override outside of the Dockerfile)
CMD [ "helloWorld" ]