# Stage 1: Build the React application
FROM node:latest AS builder

# This sets the working directory inside the container to /app.
# All subsequent commands will run from this directory.
WORKDIR /app

# Clone your specific branch from your GitHub repository
# The code will be cloned into /app/pre-cimulation
RUN git clone https://github.com/juanfesba/pre-cimulation.git .

# Change the working directory to your frontend subdirectory within the cloned repo
# We navigate into the cloned repo folder, then into the frontend folder
WORKDIR /app/precimulation-fe/frontend

# This command runs npm to install all the project dependencies
# It will use the package.json in the current working directory (/app/pre-cimulation/precimulation-fe/frontend)
RUN npm install

# This executes the build script defined in your package.json
# It will build the app in the current directory, creating ./build
RUN npm run build

# --- Inspection Steps ---
# List the contents of the current directory after the build
## RUN ls -l
# List the contents of the expected build directory
## RUN ls -l build
# If ls build fails, try listing the contents of /app to see the overall structure
# RUN ls -l /app
# --- End Inspection Steps ---

# Stage 2: Serve the built application with a lightweight web server (like Nginx)
FROM nginx:alpine

# Copy the built files from the builder stage
# The build output will be located at /app/pre-cimulation/precimulation-fe/frontend/build
COPY --from=builder /app/precimulation-fe/frontend/dist /usr/share/nginx/html

# Nginx typically serves web content on port 80
EXPOSE 80

# This specifies the command to run when the container starts
CMD ["nginx", "-g", "daemon off;"]