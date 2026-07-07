"use client"
import Navbar from "@/components/navbar";
import { useState } from "react";

export default function Register() {
  const [admnno, setAdmnno] = useState("");
  const [name, setName] = useState("");
  const [social, setSocial] = useState("");
  const [socialtype, setSocialtype] = useState("Discord");
  const [roomno, setRoomno] = useState("");
  const [blockno, setBlockno] = useState("");
  const [error, setError] = useState(""); // track error message

  const handleRegister = async () => {
    const encoder = new TextEncoder();
    const data = encoder.encode(admnno); // concat admnno + name
    const hashBuffer = await crypto.subtle.digest("SHA-256", data);
    const hashArray = Array.from(new Uint8Array(hashBuffer));
    const hashHex = hashArray.map(b => b.toString(16).padStart(2, "0")).join("");

    const queryParams = new URLSearchParams({
      admn_hash: hashHex,
      name,
      social,
      socialtype,
      roomno,
      blockno,
      created_at: "now"
    });

    const url = `http://localhost:8080/registration?${queryParams}`;
    console.log("Request URL:", url);

    fetch(url, {
      credentials: "include"
    })
      .then(response => response.text())
      .then(data => {
        console.log("Validity:", data);
        if (data === "err") {
          setError("Duplicate admission hash, someone has already logged in with your admission number");
        } else {
          setError(""); // clear error
          window.location.href = "/";
        }
      })
      .catch(err => {
        console.error("Fetch error:", err);
        setError("Server error. Please try again later.");
      });
  };

  return (
    <div>
      <Navbar />
      <div className="flex flex-col m-5 h-11/12 justify-center min-h-screen items-center">
        <fieldset className="fieldset bg-white-200 border-blue-300 rounded-box w-xs border p-4">
          <legend className="fieldset-legend text-black">Register</legend>

          <label className="label">Admission Number *</label>
          <input onChange={(e) => setAdmnno(e.target.value)} type="text" className="input bg-white border-2 border-blue-700 outline-blue-700" placeholder="Admn no." />

          <label className="label">First name (Password) *</label>
          <input onChange={(e) => setName(e.target.value)} type="text" className="input bg-white border-2 border-blue-700 outline-blue-700" placeholder="First name" />

          <br />
          <select onChange={(e) => setSocialtype(e.target.value)} value={socialtype} className="select bg-white text-black">
            <option>Discord</option>
            <option>Instagram</option>
            <option>WhatsApp</option>
          </select>

          <label className="label">Social</label>
          <input onChange={(e) => setSocial(e.target.value)} type="text" className="input bg-white border-2 border-blue-700 outline-blue-700" placeholder="User/Number" />

          <br />
          <label className="label">Block No.</label>
          <input onChange={(e) => setBlockno(e.target.value)} type="text" className="input bg-white border-2 border-blue-700 outline-blue-700" placeholder="Block No." />

          <label className="label">Room No.</label>
          <input onChange={(e) => setRoomno(e.target.value)} type="text" className="input bg-white border-2 border-blue-700 outline-blue-700" placeholder="Room No." />

          <button onClick={handleRegister} className="btn btn-info mt-4">Register</button>
        </fieldset>

        {error && (
          <div role="alert" className="alert alert-error mt-10">
            <svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6 shrink-0 stroke-current" fill="none" viewBox="0 0 24 24">
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
