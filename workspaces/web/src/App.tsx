import { useEffect, useState } from "react";
import "./App.css";
import {
  BrowserRouter,
  Navigate,
  Route,
  Routes,
  useParams,
} from "react-router-dom";

const apiUrl = import.meta.env.VITE_API_URL;

const Prediction: React.FC<{ prediction: any }> = (p) => {
  return (
    <div>
      <div>prediction</div>
      <div>
        <div>predictionId: {p.prediction.predictionId}</div>
        <div>condition: {p.prediction.condition}</div>
        <div>createdAt: {p.prediction.createdAt}</div>
        <div>concencus: </div>
      </div>
    </div>
  );
};

const User: React.FC = () => {
  const p = useParams();

  const [res, setRes] = useState<any | null>(null);

  useEffect(() => {
    fetch(apiUrl + "/user", {
      method: "POST",
      body: JSON.stringify({
        userId: p.userId,
      }),
    })
      .then((res) => res.json())
      .then((json) => setRes(json));
  }, []);

  if (res) {
    return (
      <div>
        <div>
          <div>user</div>
          <div>
            <div>avatar: {res.user.avatar}</div>
            <div>username: {res.user.id}</div>
            <div>discriminator: {res.user.discriminator}</div>
            <div>display_name: {res.user.display_name}</div>
            <div>global_name: {res.user.global_name}</div>
          </div>
        </div>
        <div>
          {res.predictions.map((e: any) => {
            return <Prediction prediction={e} key={e.predictionId} />;
          })}
        </div>
      </div>
    );
  }

  return (
    <div>
      <div>loading...</div>
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
