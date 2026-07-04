"use client"
import { useEffect, useState } from "react";
export default function Navbar() {
  const [valid, setValid] = useState(false);

  useEffect(() => {
    fetch("http://localhost:8080/verify", { credentials: "include" })
      .then(res => res.text())
      .then(data => setValid(data === "valid"))
      .catch(() => setValid(false));
  }, []);

  return (
    <div className="z-10 fixed navbar border-2 bg-blue-700 border-b-blue-700 flex justify-between">
      <a className="btn btn-ghost text-xl font-extrabold text-white" href="/">Roommate Finder</a>
      <div className="flex items-center gap-4">
        {!valid && (
          <>
            <a href="/login" className="btn btn-primary">Login</a>
          </>
        )}
        {valid && (
          <span className="text-white font-bold">Welcome back!</span>
        )}
      </div>
    </div>
  );
}
