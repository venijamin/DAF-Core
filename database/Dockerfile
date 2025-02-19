# Use the official PostgreSQL image from the Docker Hub
FROM postgres:17.3

# Set environment variables for PostgreSQL
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=mysecretpassword
ENV POSTGRES_DB=postgres

# Copy the initialization script to the Docker image
COPY init.sql /docker-entrypoint-initdb.d/

# Expose the PostgreSQL port
EXPOSE 5432