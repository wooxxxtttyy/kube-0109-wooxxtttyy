from flask import Flask

# Create a Flask application instance
app = Flask(__name__)

# Define a route for the home page ("/")
@app.route("/")
def hello_world():
   return "<p>Hello, World!</p>"

# Define a route with a dynamic part (e.g., a name)
@app.route("/greet/<name>")
def greet_user(name):
   return "<h1>Hello!</h1>"

# Run the application if this script is executed directly
if __name__ == "__main__":
   app.run(debug=True, host="0.0.0.0") # debug=True allows for automatic reloading on code changes