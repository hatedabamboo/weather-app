import requests
from flask import Flask, jsonify, request

app = Flask(__name__)

WEATHER_API_KEY = "YOUR_WEATHER_API_KEY"


def get_location(ip):
    try:
        response = requests.get(f"https://ipapi.co/{ip}/json/")
        if response.status_code == 200:
            data = response.json()
            print(data.get("city"), data.get("latitude"), data.get("longitude"))
            return data.get("city"), data.get("latitude"), data.get("longitude")
        return None, None, None
    except Exception as e:
        print(f"Error fetching location: {e}")
        return None, None, None


def get_weather(lat, lon):
    try:
        url = f"http://api.openweathermap.org/data/2.5/weather?lat={lat}&lon={lon}&appid={WEATHER_API_KEY}&units=metric"
        response = requests.get(url)
        print(response)
        if response.status_code == 200:
            data = response.json()
            weather_desc = data["weather"][0]["description"]
            temperature = data["main"]["temp"]
            return f"Weather: {weather_desc}, Temperature: {temperature}°C"
        return "Unable to fetch weather data"
    except Exception as e:
        print(f"Error fetching weather: {e}")
        return "Error fetching weather"


@app.route("/", methods=["GET"])
def index():
    client_ip = request.remote_addr
    city, lat, lon = get_location(client_ip)

    if city and lat and lon:
        # Get weather information
        weather_info = get_weather(lat, lon)
        print(weather_info)
        return jsonify({"ip": client_ip, "city": city, "weather": weather_info})
    else:
        return jsonify({"error": "Unable to determine location from IP address"}), 500


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8080)
