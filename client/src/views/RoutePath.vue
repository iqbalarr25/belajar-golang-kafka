<template>
  <div id="app">
    <div :id="mapId" class="map"></div>
    <LRoutingMachine :mapObject="mapObject" :waypoints="waypoints"/>
  </div>
</template>

<script>
import "leaflet-routing-machine/dist/leaflet-routing-machine.css";
import "leaflet/dist/leaflet.css";
import LRoutingMachine from "@/components/LRoutingMachine.vue";
import L from "leaflet";

const waypoints = [
  {lat: -6.314815791435494, lng: 106.90667133736513},
  {lat: -6.299168137712657, lng: 106.89046710733886},
];

export default {
  components: {
    LRoutingMachine,
  },
  data() {
    return {
      mapId: "map",
      mapObject: null,
      zoom: 6,
      center: {lat: 38.7436056, lng: -2.2304153},
      osmUrl: "http://{s}.tile.osm.org/{z}/{x}/{y}.png",
      attribution:
          '&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors',
      waypoints,
    };
  },
  mounted() {
    this.mapObject = L.map(this.mapId, {
      zoom: this.zoom,
      center: this.center,
    });

    var redIcon = new L.Icon({
      iconUrl: 'https://cdn.rawgit.com/pointhi/leaflet-color-markers/master/img/marker-icon-2x-red.png',
      shadowUrl: 'https://cdnjs.cloudflare.com/ajax/libs/leaflet/1.3.4/images/marker-shadow.png',
      iconSize: [25, 41],
      iconAnchor: [12, 41],
      popupAnchor: [1, -34],
      shadowSize: [41, 41],
    });

    const blueIcon = new L.Icon({
      iconUrl: 'https://cdn.rawgit.com/pointhi/leaflet-color-markers/master/img/marker-icon-2x-blue.png',
      shadowUrl: 'https://cdnjs.cloudflare.com/ajax/libs/leaflet/1.3.4/images/marker-shadow.png',
      iconSize: [25, 41],
      iconAnchor: [12, 41],
      popupAnchor: [1, -34],
      shadowSize: [41, 41]
    });

    L.marker([-6.314815791435494, 106.90667133736513], {
      icon: blueIcon,
      title:"origin",
    }).addTo(this.mapObject)

    L.marker([-6.299168137712657, 106.89046710733886], {
      icon: redIcon,
      title:"destination",
    }).addTo(this.mapObject)

    L.tileLayer(this.osmUrl, {
      attribution: this.attribution,
    }).addTo(this.mapObject);
  },
};
</script>

<style>
html,
body,
#app {
  height: 100%;
  margin: 0;
}

.map {
  position: absolute;
  width: 100%;
  height: 100%;
}
</style>
