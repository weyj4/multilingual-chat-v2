import React, { Component } from 'react'
import { observer } from 'mobx-react'
import store from './messageStore'
import MessageInput from './MessageInput'

@observer
export default class Messages extends Component {
  render() {
    return (
      <div>
        {store.userMessages.map((msg, i) => <p key={i}>{msg}</p>)}
        <MessageInput />
        <button onClick={store.initializeSocket} disabled={store.loadingWS}>Enter Chat</button>
        {store.allMessages.map((msg, i) => <p key={i}>{msg}</p>)}
      </div>
    )
  }
}
