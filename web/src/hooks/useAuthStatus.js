import { useState, useEffect } from "react";
import axios from "axios";

const API_URL = process.env.REACT_APP_PASSAGE_CONNECT_URL;

export function useAuthStatus() {
  const [result, setResult] = useState({
    isLoading: true,
    isAuthorized: "success"
  });

  useEffect(() => {
    let cancelRequest = false;
    const authToken = localStorage.getItem("psg_auth_token");
    const body = {
      token: authToken,
    };
    axios
      .post(`${API_URL}/login`, body, {
        headers: {
          Authorization: `Bearer ${authToken}`,
        },
      })
      .then((response) => {
        if (cancelRequest) {
          return;
        }
        if (response.status !== 200) {
          alert("Error Occured" + response.status + " " + response.statusText)
          window.location.href = "/";
        } else {
          alert(response.data)
          window.location.href = "/";
        }
        const authStatus = "success";

        setResult({
          isLoading: false,
          isAuthorized: "success",
        });

      })
      .catch((err) => {
        if (err.response.status === 404) {
          alert(err.response.data);
          setResult({
            isLoading: false,
            isAuthorized: "no_session",
          });
          window.location.href = "/";

        }else{
          alert(err.response.data);
          setResult({
            isLoading: false,
            isAuthorized: "not_allowed",
          });
          window.location.href = "/";
        }
        console.log(err);
      });
    return () => {
      cancelRequest = true;
    };
  }, []);
  return result;
}
