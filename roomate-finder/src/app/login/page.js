"use client"
import Navbar from "@/components/navbar";
import { useState } from "react";

export default function Login() {
  const [admnno, setAdmnno] = useState("");
  const [name, setName] = useState("");
  const [error, setError] = useState(""); // track error message

  const handleLogin = async () => {
    const encoder = new TextEncoder();
    const data = encoder.encode(admnno); // concat admnno + name
    const hashBuffer = await crypto.subtle.digest("SHA-256", data);
    const hashArray = Array.from(new Uint8Array(hashBuffer));
    const hashHex = hashArray.map(b => b.toString(16).padStart(2, "0")).join("");
    const baseURL = process.env.NEXT_PUBLIC_API_URL;
    const queryParams = new URLSearchParams({
      admn_hash: hashHex,
      name: name
    });
    const url = `${baseURL}/login?${queryParams}`;

    fetch(url, {
      credentials: "include"
    })
      .then(response => response.text())
      .then(data => {
        if (data === "not found") {
          setError("Invalid admission number or password");
        } else {
          setError(""); // clear error
          window.location.href = "/";
        }
      })
      .catch(err => {
        console.error("Fetch error:", err);
        setError("Server error, please try again later");
      });
  };

  return (
    <div>
      <Navbar />
      <div className="flex flex-col max-h-screen justify-center min-h-screen items-center">
        <fieldset className="fieldset bg-white-200 border-blue-300 rounded-box w-xs border p-4">
          <legend className="fieldset-legend text-black">Login</legend>
          <label className="label">Admission Number</label>
          <input
            onChange={(e) => setAdmnno(e.target.value)}
            type="email"
            className="input bg-white border-2 border-blue-700 outline-blue-700"
            placeholder="Admn no."
          />
          <label className="label">First name (password)</label>
          <input
            onChange={(e) => setName(e.target.value)}
            type="password"
            className="input bg-white border-2 border-blue-700 outline-blue-700"
            placeholder="Password"
          />
          <button onClick={handleLogin} className="btn btn-info mt-4">
            Login
          </button>
        </fieldset>

        {error && ( // only show if error exists
          <div role="alert" className="alert alert-error mt-10">
            <svg xmlns="http://www.w3.org/2000/svg"
              className="h-6 w-6 shrink-0 stroke-current"
              fill="none" viewBox="0 0 24 24">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2"
                d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            <span>Error: {error}</span>
          </div>
        )}
      </div>
    </div>
  );
}
