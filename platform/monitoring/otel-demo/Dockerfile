# Use an official Maven image to build the application
FROM maven:3.8.6-eclipse-temurin-17 AS build

# Set the working directory in the container
WORKDIR /app

# Copy the Maven project files (pom.xml and source code)
COPY pom.xml .
COPY src ./src
COPY opentelemetry-javaagent.jar .

# Build the application
RUN mvn clean package -X -DskipTests

# Use a smaller JDK image to run the application
FROM eclipse-temurin:17-jdk

# Set the working directory in the container
WORKDIR /app

# Copy the built JAR file from the previous stage
COPY --from=build /app/target/*.jar app.jar
COPY --from=build /app/opentelemetry-javaagent.jar .

# Expose the application's port
EXPOSE 8080

# Run the application
CMD ["java", "-javaagent:opentelemetry-javaagent.jar", "-jar", "app.jar"]