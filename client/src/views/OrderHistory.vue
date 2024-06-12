<template>
  <MapView ref="mapView" :icon="icon" :routes="routes"/>
</template>

<script>
import MapView from '../components/MapView.vue'
import {useRoute} from 'vue-router'
import axios from "axios";
import 'leaflet.animatedmarker/src/AnimatedMarker'

export default {
  name: 'LiveTracking',
  components: {
    MapView
  },
  data() {
    return {
      id: null,
      icon: 'https://cdn0.iconfinder.com/data/icons/small-n-flat/24/678111-map-marker-512.png',
      socket: null,
      routes: [],
      center: [-0.789275, 113.921327],
      zoom: 5,
    }
  },
  mounted() {
    const route = useRoute()
    const apiKey = route.query.api_key
    const externalIdentity = route.query.external_identity
    let url = "http://localhost:9000/api/v1/orders/history";
    if (process.env.NODE_ENV === 'production') {
      url = "https://" + window.location.host + "/api/v1/orders/history";
    }
    axios.post(url, {
      api_key: apiKey,
      external_identity: externalIdentity
    }, {
      headers: {
        'Content-Type': 'application/json',
      }
    }).then((response) => {
      this.routes = response.data.data
      if (this.routes.length > 0) {
        this.$refs.mapView.updateCenter(response.data.data[0])
        this.$refs.mapView.initAnimated(this.routes)
      }
    }).then(() => {
      let url_destination = "http://localhost:9000/api/v1/orders/view";
      if (process.env.NODE_ENV === 'production') {
        url_destination = "https://" + window.location.host + "/api/v1/orders/view";
      }
      axios.post(url_destination, {
        api_key: apiKey,
        external_identity: externalIdentity
      }, {
        headers: {
          'Content-Type': 'application/json',
        }
      }).then((response) => {
        const pickUp = response?.data?.data?.filter(item => item.type === "PICKUP") ?? []
        const dropOff = response?.data?.data?.filter(item => item.type === "DROPOFF") ?? []
        if (pickUp.length > 0) this.$refs.mapView.updateMarkerPickUp([pickUp[0].lat, pickUp[0].lng])
        if (dropOff.length > 0) this.$refs.mapView.updateMarkerDropOff([dropOff[0].lat, dropOff[0].lng])
      }).catch((error) => {
        console.log(error)
      })
    }).catch((error) => {
      console.log(error)
    })
  },
}
</script>

<style>
body {
  min-height: 100vh;
  margin: 0;
}
</style>
