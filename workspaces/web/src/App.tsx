import { useEffect, useState } from "react";
import {
  BrowserRouter,
  Navigate,
  Route,
  Routes,
  useParams,
  useNavigate,
} from "react-router-dom";

const apiUrl = import.meta.env.VITE_API_URL;

const Prediction: React.FC<{ prediction: any }> = (p) => {
  const voters = p.prediction.voters;

  const percentage = (): number => {
    const p = voters.filter((e: any) => e.vote).length / voters.length;
    return p * 100;
  };

  return (
    <div>
      <h2>
        {p.prediction.predictionId}
        {voters.length > 0 ? ` (${percentage()}%)` : ""}
      </h2>
      <div>condition(s): {p.prediction.condition}</div>
      <div>made on: {p.prediction.createdAt}</div>
    </div>
  );
};

const User: React.FC = () => {
  const nav = useNavigate();
  const p = useParams();

  const [res, setRes] = useState<any | null>(null);

  useEffect(() => {
    fetch(apiUrl + "/user", {
      method: "POST",
      body: JSON.stringify({
        userId: p.userId,
      }),
    })
      .then((res) => {
        if (!res.ok) throw new Error("no");
        return res;
      })
      .then((res) => res.json())
      .then((json) => setRes(json))
      .catch(() => {
        nav("/error");
      });
  }, []);

  if (res) {
    const user = res.user;

    return (
      <div>
        <div>
          {/* todo: avatar */}
          {/* todo: display/global name */}
          <h1>
            {user.username}#{user.discriminator}'s predictos
          </h1>
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
