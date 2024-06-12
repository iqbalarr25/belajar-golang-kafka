<template>
  <div style="height:100vh; width:100%">
    <l-map ref="baseMap" :zoom="zoom" :center="center" @ready="onReady">
      <l-tile-layer :url="url" :subdomains="subdomains" :attribution="attribution"></l-tile-layer>
      <!--      <l-marker :lat-lng="markerLatLng" :visible="markerVisible">-->
      <!--        <l-icon :icon-url="icon" :icon-size="[25, 41]" :icon-anchor="[12, 41]"></l-icon>-->
      <!--      </l-marker>-->
      <l-marker :lat-lng="markerPickUpLatLng" :visible="markerPickUpVisible">
        <l-icon :icon-url="pickupIcon" :icon-size="[25, 41]"></l-icon>
        <l-tooltip>Pick Up</l-tooltip>
      </l-marker>
      <l-marker :lat-lng="markerDropOffLatLng" :visible="markerDropOffVisible">
        <l-icon :icon-url="dropOffIcon" :icon-size="[25, 41]" :icon-anchor="[12, 41]"></l-icon>
        <l-tooltip>Drop Off</l-tooltip>
      </l-marker>
      <l-polyline :lat-lngs="routes" :color="color"></l-polyline>
    </l-map>
  </div>
</template>

<script>
import "leaflet/dist/leaflet.css";
import L from "leaflet";
import {LIcon, LMap, LMarker, LPolyline, LTileLayer, LTooltip} from "@vue-leaflet/vue-leaflet";
import 'leaflet.animatedmarker/src/AnimatedMarker'
import '../assets/leaflet.rotatedMarker'

export default {
  components: {
    LTooltip,
    LPolyline,
    LIcon,
    LMarker,
    LMap,
    LTileLayer,
  },
  props: {
    icon: {
      type: String,
      default: () => 'https://cdn0.iconfinder.com/data/icons/small-n-flat/24/678111-map-marker-512.png',
    },
    pickupIcon: {
      type: String,
      default: () => 'https://cdn.rawgit.com/pointhi/leaflet-color-markers/master/img/marker-icon-2x-green.png',
    },
    dropOffIcon: {
      type: String,
      default: () => 'https://cdn.rawgit.com/pointhi/leaflet-color-markers/master/img/marker-icon-2x-red.png',
    },
    routes: {
      type: Array,
      default: () => [],
    },
    color: {
      type: String,
      default: () => 'green',
    },
  },
  data() {
    return {
      isMapReady: false,
      isFirstLoad: true,
      routingLayer: null,
      animatedMarker: null,
      // waypoints: [],
      waypoints: [{lat: -0.789275, lng: 113.921327}, {lat: -0.789275, lng: 113.921327}],
      // waypoints: [{lat: -2.520773, lng: 104.129249}, {lat: 1.067563, lng: 101.337576}],
      url: 'https://{s}.google.com/vt/lyrs=m&x={x}&y={y}&z={z}',
      subdomains: ['mt0', 'mt1', 'mt2', 'mt3'],
      attribution: '&copy; <a href="https://maps.google.com" target="_blank">Google Map</a> | <a href="https://transtrack.co">TransTRACK</a>',
      center: [-0.789275, 113.921327],
      zoom: 5,
      marker: null,
      markerLatLng: [-0.789275, 113.921327],
      markerPickUpLatLng: [-0.789275, 113.921327],
      markerDropOffLatLng: [-0.789275, 113.921327],
      markerVisible: false,
      markerPickUpVisible: false,
      markerDropOffVisible: false,
    };
  },
  methods: {
    onReady() {
      console.log('Map is ready')
      this.isMapReady = true
    },
    updateMarker(latLng, rotate = 0) {
      this.markerVisible = true
      this.center = latLng
      this.zoom = 15
      if (this.isFirstLoad) {
        this.markerLatLng = latLng
        this.isFirstLoad = false
        this.marker = L.marker(latLng, {
          rotationAngle: rotate,
          icon: L.icon({
            iconUrl: this.icon,
            iconSize: [25, 41],
            iconAnchor: [16, 37]
          }),
        }).addTo(this.$refs.baseMap.leafletObject);
      } else {
        const start = this.markerLatLng
        const end = latLng
        const duration = 30000
        const steps = 100
        const stepLatLng = [(end[0] - start[0]) / steps, (end[1] - start[1]) / steps]

        if (stepLatLng[0] !== 0 && stepLatLng[1] !== 0) {
          let step = 0
          const interval = setInterval(() => {
            this.markerLatLng = [start[0] + stepLatLng[0] * step, start[1] + stepLatLng[1] * step]
            this.marker.setRotationAngle(rotate)
            this.marker.setLatLng(new L.latLng(this.markerLatLng))
            step++
            if (step >= steps) {
              clearInterval(interval)
            }
          }, duration / steps)
        }
      }
    },
    updateCenter(latLng) {
      this.center = latLng
      this.zoom = 17
    },
    updateMarkerPickUp(latLng, visible = true) {
      this.markerPickUpVisible = visible
      this.markerPickUpLatLng = latLng
    },
    updateMarkerDropOff(latLng, visible = true) {
      this.markerDropOffVisible = visible
      this.markerDropOffLatLng = latLng
      this.waypoints[1] = {lat: latLng[0], lng: latLng[1]}
    },
    initRoute(waypoints) {
      this.routingLayer = new L.Routing.control({
        lineOptions: {
          addWaypoints: false,
          styles: [{color: '#198754', opacity: 1, weight: 5}]
        },
        draggableWaypoints: false,
        routeWhileDragging: false,
        createMarker: function () {
          return null;
        },
        waypoints: waypoints
      });
      this.routingLayer.addTo(this.$refs.baseMap.leafletObject);
    },
    updateRoute(position) {
      this.routingLayer.setWaypoints([
        L.latLng(position[0], position[1]),
        this.routingLayer.options.waypoints[1]
      ])
    },
    initAnimated(routes) {
      const line = L.polyline(routes)
      this.animatedMarker = L.animatedMarker(line.getLatLngs(), {
        // distance: 300,  // meters
        // interval: 2000, // milliseconds
        rotationAngle: 0,
        icon: L.icon({
          iconUrl: "https://telematics.transtrack.id/images/device_icons/625e423669a233.80559840_online.png",
          iconSize: [25, 41],
          iconAnchor: [16, 37]
        }),
      });
      this.zoom = 13
      setTimeout(() => {
        this.$refs.baseMap.leafletObject.addLayer(this.animatedMarker);
      }, 1000)
    }
  },
};
</script>

<style scoped>

</style>