// this is for everything about the chat room
const newMessage = new Vue({
  el: '#send-message',
  data: {
    newMessageText: ''
  },
  methods: {
    submitNewMessage: function () {
      const msg = messageFactor(this.newMessageText)
      conn.send(JSON.stringify(msg))
      this.newMessageText = ''
    }
  }
})
