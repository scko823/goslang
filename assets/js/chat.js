// this is for everything about the chat room
conn = new WebSocket("ws://" + document.location.host + "/ws?room=" + (getParameterByName('room')||'main') )

const messageList = Vue.component('message-list', {
  template: `<div id="message">
    <ul id="message-list">
      <li v-for="msg in messages">
        {{msg.message}}
      </li>
    </ul>
  </div>`,
  data: function () {
    return {
      messages: []
    }
  },
  mounted: function () {
    const self = this
    conn.onmessage = function(evt){
      let message = JSON.parse(event.data)
      self.messages = self.messages.concat([message])
    }
  }
})


Vue.component('send-message',{
  template:`<div id="send-message">
    <form id="sumbit" class="message" @submit.prevent="submitNewMessage">
      <input v-model="newMessageText" />
    </form>
  </div>`,
  data: function () {
    return {
      newMessageText: ''
    }
  },
  methods: {
    submitNewMessage: function () {
      const msg = messageFactor(this.newMessageText)
      conn.send(JSON.stringify(msg))
      this.newMessageText = ''
    }
  }
})

new Vue({
  el:'#app'
})
