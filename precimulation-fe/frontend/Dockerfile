# Stage 1: Build the React application
FROM node:latest AS builder

# This sets the working directory inside the container to /app.
# All subsequent commands will run from this directory.
WORKDIR /app

# This copies your package.json and package-lock.json (or yarn.lock) files
# into the /app directory. Copying these first allows Docker to cache the
# npm install step if your dependencies haven't changed, speeding up subsequent builds.
# COPY package*.json ./
# BUT NOW WE DO IT WITH GIT
RUN git clone https://github.com/juanfesba/pre-cimulation.git .
WORKDIR /app/precimulation-fe/frontend

# This command runs npm to install all the project dependencies based on the copied package*.json files.
RUN npm install

# This executes the build script defined in your package.json (usually react-scripts build or similar).
# This compiles your React application and outputs the production-ready static files into a build folder
# (or similar, depending on your setup) within the working directory.
RUN npm run build

# Stage 2: Serve the built application with a lightweight web server (like Nginx)
# It's also based on Alpine Linux. Nginx is a popular web server that's great for
# serving static files. This stage does not include Node.js or your development dependencies.
FROM nginx:alpine

# This is the key to the multi-stage build. It copies the built static files
# from the build directory in the builder stage (the first stage) to the
# default Nginx webroot directory (/usr/share/nginx/html) in the current stage.
COPY --from=builder /app/precimulation-fe/frontend/build /usr/share/nginx/html

# Nginx typically serves web content on this port by default.
# Curious that vite uses 5173 by default, but Nginx uses 80.
EXPOSE 80

# This specifies the command to run when the container starts. It tells Nginx to
# run in the foreground, which is necessary for it to keep the container alive.
CMD ["nginx", "-g", "daemon off;"]