import './index.css'
import {BrowserRouter as Router, Route, Routes} from "react-router-dom";
import CreateUser from "./CreateUser";
import MainChat from "./MainChat.jsx";
import Login from "./Login.jsx";

function App() {

  return (
      <Router>
          <Routes>
              <Route path="/create-user" element={<CreateUser />} />
              <Route path="/chat" element={<MainChat />} />
              <Route path="/chat/:channelId" element={<MainChat />} />
              <Route path="/" element={<Login />} />
          </Routes>
      </Router>
  )
}

export default App
