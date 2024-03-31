# openmeteo exporter

Simple prometheus exporter, that fetches weather for a location using [Openmeteo's])[https://open-meteo.com] API.

```
# HELP openmeteo_apparent_temperature apparent temperature, °C
# TYPE openmeteo_apparent_temperature gauge
openmeteo_apparent_temperature{location="home"} 19.1
# HELP openmeteo_cloud_cover cloud cover, %
# TYPE openmeteo_cloud_cover gauge
openmeteo_cloud_cover{location="home"} 100
# HELP openmeteo_interval interval of report, seconds
# TYPE openmeteo_interval gauge
openmeteo_interval{location="home"} 900
# HELP openmeteo_is_day Is daytime, 1 when daytime, 0 when not
# TYPE openmeteo_is_day gauge
openmeteo_is_day{location="home"} 1
# HELP openmeteo_precipitation precipitation, mm
# TYPE openmeteo_precipitation gauge
openmeteo_precipitation{location="home"} 0
# HELP openmeteo_pressure_msl pressure at MSL, hPa
# TYPE openmeteo_pressure_msl gauge
openmeteo_pressure_msl{location="home"} 996.7
# HELP openmeteo_pressure_surface pressure at surface, hPa
# TYPE openmeteo_pressure_surface gauge
openmeteo_pressure_surface{location="home"} 992.5
# HELP openmeteo_rain rain, mm
# TYPE openmeteo_rain gauge
openmeteo_rain{location="home"} 0
# HELP openmeteo_relative_humidity relative humidity (2m), %
# TYPE openmeteo_relative_humidity gauge
openmeteo_relative_humidity{location="home"} 58
# HELP openmeteo_showers showers, mm
# TYPE openmeteo_showers gauge
openmeteo_showers{location="home"} 0
# HELP openmeteo_snowfall snowfall, cm
# TYPE openmeteo_snowfall gauge
openmeteo_snowfall{location="home"} 0
# HELP openmeteo_temperature current temperature (2m), °C
# TYPE openmeteo_temperature gauge
openmeteo_temperature{location="home"} 19
# HELP openmeteo_time Time of observation, unix format
# TYPE openmeteo_time gauge
openmeteo_time{location="home"} 1.7118918e+09
# HELP openmeteo_weather_code ???, wmo code
# TYPE openmeteo_weather_code gauge
openmeteo_weather_code{location="home"} 3
# HELP openmeteo_wind_direction direction of wind (10m), °
# TYPE openmeteo_wind_direction gauge
openmeteo_wind_direction{location="home"} 270
# HELP openmeteo_wind_gusts wind gusts (10m), km/h
# TYPE openmeteo_wind_gusts gauge
openmeteo_wind_gusts{location="home"} 13
# HELP openmeteo_wind_speed wind speed (10m), km/h
# TYPE openmeteo_wind_speed gauge
openmeteo_wind_speed{location="home"} 1.1
```
