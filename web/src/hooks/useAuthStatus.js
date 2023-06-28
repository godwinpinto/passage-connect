import { useState, useEffect } from "react";
import axios from "axios";

const API_URL = process.env.REACT_APP_PASSAGE_CONNECT_URL;

export function useAuthStatus() {
  const [result, setResult] = useState({
    isLoading: true,
    isAuthorized: false,
    username: "",
  });

  useEffect(() => {
    let cancelRequest = false;
    const authToken = localStorage.getItem("psg_auth_token");
    const body = {
      token:authToken,
    };
    axios
      .post(`${API_URL}/login`, null, {
        headers: {
          Authorization: `Bearer ${authToken}`,
        },
      })
      .then((response) => {
        if (cancelRequest) {
          return;
        }
        if(response.status !== 200){
          alert("Error Occured"+response.status+" "+response.statusText)
          window.location.href = "/";
        }else{
          alert("You may now proceed with your SSH")
          window.location.href = "/";
        }
        const authStatus = "success";

        if (authStatus === "success") {
          setResult({
            isLoading: false,
            isAuthorized: authStatus,
            username: "Log in successful",
          });
        } else {
          setResult({
            isLoading: false,
            isAuthorized: false,
            username: "",
          });
        }
      })
      .catch((err) => {
        console.log(err);
        setResult({
          isLoading: false,
          isAuthorized: false,
          username: "",
        });
      });
    return () => {
      cancelRequest = true;
    };
  }, []);
  return result;
}
