<template>
  <MapView ref="mapView" :icon="icon"/>
</template>

<script>
import MapView from '../components/MapView.vue'
import {useRoute} from 'vue-router'
import axios from "axios";

export default {
  name: 'LiveTracking',
  components: {
    MapView
  },
  data() {
    return {
      id: null,
      icon: 'https://cdn0.iconfinder.com/data/icons/small-n-flat/24/678111-map-marker-512.png',
      socket: null
    }
  },
  beforeMount() {
    const route = useRoute()
    const apiKey = route.query.api_key
    const externalIdentity = route.query.external_identity
    this.getMark(apiKey, externalIdentity)
  },
  mounted() {
    const route = useRoute()
    const apiKey = route.query.api_key
    const externalIdentity = route.query.external_identity

    let url = "http://localhost:9000/api/v1/live-locations";
    if (process.env.NODE_ENV === 'production') {
      url = "https://" + window.location.host + "/api/v1/live-locations";
    }
    axios.get(url, {
      params: {
        api_key: apiKey,
        external_identity: externalIdentity
      },
      headers: {
        'Content-Type': 'application/json',
      }
    }).then((response) => {
      this.icon = response.data.data.icon
      if (this.$refs.mapView.isMapReady) {
        this.$refs.mapView.updateMarker([response.data.data.lat, response.data.data.lng],response.data.data.course)
      } else {
        setTimeout(() => {
          this.$refs.mapView.updateMarker([response.data.data.lat, response.data.data.lng],response.data.data.course)
        }, 4000)
      }
    }).catch((error) => {
      console.log(error)
    })

    let ws = "ws://localhost:9000/ws/live-locations";
    if (process.env.NODE_ENV === 'production') {
      ws = "wss://" + window.location.host + "/ws/live-locations";
    }
    this.websocket = new WebSocket(ws + "?api_key=" + apiKey + "&external_identity=" + externalIdentity)
    this.websocket.onmessage = (event) => {
      const data = JSON.parse(event.data)
      this.icon = data.icon
      if (this.$refs.mapView?.isMapReady) {
        this.$refs.mapView.updateMarker([data.lat, data.lng],data.course)
        this.$refs.mapView.updateRoute([data.lat, data.lng],data.course)
      } else {
        setTimeout(() => {
          this.$refs.mapView.updateMarker([data.lat, data.lng])
          this.$refs.mapView.updateRoute([data.lat, data.lng])
        }, 1500)
      }
      console.log(data)
    }
  },
  methods: {
    getMark(apiKey, externalIdentity) {
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

        if (pickUp.length > 0 && dropOff.length > 0) {
          console.log("true")
          // this.$refs.mapView.updateMarkerPickUp([pickUp[0].lat, pickUp[0].lng])
          this.$refs.mapView.updateMarkerDropOff([dropOff[0].lat, dropOff[0].lng])
          this.$refs.mapView.initRoute([{lat: pickUp[0].lat, lng: pickUp[0].lng}, {lat: dropOff[0].lat, lng: dropOff[0].lng}])
        }
      }).catch((error) => {
        console.log(error)
      })
    }
  }
}
</script>

<style>
body {
  min-height: 100vh;
  margin: 0;
}
</style>
