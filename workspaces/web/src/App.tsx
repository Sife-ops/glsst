import "./App.css";
import { BrowserRouter, Navigate, Route, Routes } from "react-router-dom";

const User: React.FC = () => {
  return (
    <div>
      <div>user</div>
    </div>
  );
};

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/user/:userId" element={<User />} />
        <Route path="/error" element={<div>error</div>} />
        <Route path="*" element={<Navigate to="/error" />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
