/* eslint-disable sort-keys */
import util from 'util'

import {
  grpc,
  IPuppetServer,
  PuppetService,
  EventResponse,
  EventType,
  DingResponse,
  EventRequest,
}                       from '../src/mod'

// import { StringValue } from 'google-protobuf/google/protobuf/wrappers_pb'

import {
  puppetServerImpl,
}                     from '../tests/puppet-server-impl'

let eventStream: undefined | grpc.ServerWritableStream<EventRequest, EventResponse>
const dingQueue = [] as string[]

/**
 * Implements the SayHello RPC method.
 */
const puppetServerExample: IPuppetServer = {
  ...puppetServerImpl,

  event: (streammingCall) => {
    console.info('event(streamingCall)')

    if (eventStream) {
      console.info('event() end old eventStream to accept the new one.')
      eventStream.end()
      eventStream = streammingCall
    }

    eventStream = streammingCall
    while (dingQueue.length > 0) {
      const data = dingQueue.shift()
      const eventResponse = new EventResponse()
      eventResponse.setType(EventType.EVENT_TYPE_DONG)
      eventResponse.setPayload(data!)
      eventStream.write(eventResponse)
    }
    /**
      * Detect if Inexor Core is gone (GRPC disconnects)
      *  https://github.com/grpc/grpc/issues/8117#issuecomment-362198092
      */
    eventStream.on('cancelled', () => {
      console.info('eventStream.on(calcelled)')
      eventStream?.end()
      eventStream = undefined
    })
  },

  ding: (call, callback) => {
    const data = call.request.getData()
    console.info(`ding(${data})`)

    if (!eventStream) {
      dingQueue.push(data)
    } else {
      const eventResponse = new EventResponse()
      eventResponse.setType(EventType.EVENT_TYPE_DONG)
      eventResponse.setPayload(data)
      eventStream.write(eventResponse)
    }

    callback(null, new DingResponse())
  },
}

/**
 * Starts an RPC server that receives requests for the Greeter service at the
 * sample server port
 */
async function main () {
  const server = new grpc.Server()
  server.addService(
    PuppetService,
    puppetServerExample,
  )
  const port = await util.promisify(server.bindAsync.bind(server))(
    '127.0.0.1:8788',
    grpc.ServerCredentials.createInsecure(),
  )
  console.info('Listen on port:', port)
  server.start()
  return 0
}

main()
  // .then(process.exit)
  .catch(e => {
    console.error(e)
    process.exit(1)
  })
