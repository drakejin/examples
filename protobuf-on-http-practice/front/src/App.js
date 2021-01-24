import pbhttp from './proto/http_pb';
// import proto from './proto/http.proto';
// import protobuf from 'protobufjs';
import axios from 'axios'
import './App.css';
import React, { useState } from 'react';


function App() {
  const [query, setQuery] = useState("query")

  return (
    <div className="App">
      <header className="App-header">
        <p>
          protobuf 의 메세지 프로퍼티 id값을 변경하지 않고 추가만 했는데 <br/>
          추가하고 나니 javascript에서 깨져버린다구?
        </p>
        <input
          type="query"
          style={{ fontSize: '2rem' }}
          onChange={(e) => setQuery(e.target.value)}
        ></input>
        <button
          style={{ fontSize: '2rem' }}
          onClick={() => send(query)}
        >
          Send
        </button>
        <p>
          F12를 눌러 결과를 확인하세요
        </p>

      </header>
    </div>
  );
}

function send(query) {
  const req = new pbhttp.Request()

  const reqA = new pbhttp.RequestA()
  reqA.setQuery(query)
  req.setReqA(reqA)

  const reqB = new pbhttp.RequestB()
  reqB.setQuery(query)
  req.setReqB(reqB)

  const reqC = new pbhttp.RequestC()
  reqC.setQuery(query)
  req.setReqC(reqC)

  callFetch(req)
  // callAxios(req)
}

const callAxios = (req) => {
  axios.post(
    "http://0.0.0.0:4100",
    req.serializeBinary(),
    {
      headers: {
        'Content-Type': 'application/protobuf',
      },
      responseType: 'arraybuffer',
    },
  ).then((response) => {
    const bytes = response.data
    const data = pbhttp.Response.deserializeBinary(bytes).toObject()
    console.log(data)
  })
}


const callFetch = (req) => {
  fetch("http://0.0.0.0:4100", {
    method: 'POST',
    headers: {
        'Content-Type': 'application/protobuf',
    },
    body: req.serializeBinary()
  })
  .then(async (response) => {
    console.log(response)
    const bytes = await response.arrayBuffer()
    const data = pbhttp.Response.deserializeBinary(bytes).toObject()
    console.log(data)
  })
}

export default App;



/*
    protobuf.load(proto, function(err, root) {
      if (err) throw err;

      var ResponseMessage = root.lookupType("drakejin.examples.http.Response");

      console.log(ResponseMessage)

      // Exemplary payload
      var payload = pbhttp.Response.deserializeBinary(bytes).toObject()

      console.log(payload)
      var errMsg = ResponseMessage.verify(payload);
      if (errMsg) throw Error(errMsg);

      // case1
      var message = ResponseMessage.create(payload); // or use .fromObject if conversion is necessary
      console.log("case1", message.toJSON())

      // case2
      var buffer = ResponseMessage.encode(message).finish();
      console.log("case2", buffer)

      // case3
      var decoded = ResponseMessage.decode(buffer);
      console.log("case3", decoded.toJSON())


      // case4: Maybe convert the message back to a plain object
      var object = ResponseMessage.toObject(message, {
          longs: String,
          enums: String,
          bytes: String,
          // see ConversionOptions
      });
      console.log("case4", object)
    })
  });

  */
