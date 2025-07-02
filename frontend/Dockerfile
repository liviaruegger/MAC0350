# Use an official Node.js runtime as a parent image
FROM node:20

# Set the working directory in the container
WORKDIR /app

# Copy package.json and package-lock.json to the working directory
COPY package*.json ./

# Install the project dependencies
RUN npm install

# Copy the rest of the application's source code from the host to the container
COPY . .

# Make port 5173 available to the world outside this container
EXPOSE 5173

# Run the app when the container launches
CMD ["npm", "run", "dev", "--", "--host"]