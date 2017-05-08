import mobx, { observable, computed } from 'mobx'

class MessageStore {
  @observable userMessages = []
  @observable allMessages = []
  @observable loadingWS = true

  constructor() {
    this.initializeSocket = this.initializeSocket.bind(this)
  }

  connect() {
    return new Promise((resolve, reject) => {
      this.socket = new WebSocket('ws://localhost:8080/socket')
      this.socket.onopen = () => {
        this.loadingWS = false
        resolve()
      }
      this.socket.onmessage = (e) => this.allMessages.push(e.data)
    })
  }

  initializeSocket() {
    console.log('this', this)
    this.socket.send('init')
  }

  addMessage(message) {
    this.userMessages.push(message)
    this.socket.send(message)
  }
}

const store = new MessageStore()
store.connect()

export default store
