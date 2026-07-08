/* eslint-disable @next/next/no-html-link-for-pages */
/* eslint-disable react/jsx-no-comment-textnodes */
"use client"
import { useEffect, useState } from "react";

export default function Navbar() {
  const [valid, setValid] = useState(false);
  const [showConfirm, setShowConfirm] = useState(false);

  const baseURL = process.env.NEXT_PUBLIC_API_URL;

  useEffect(() => {
    fetch(`${baseURL}/verify`, { credentials: "include" })
      .then(res => res.text())
      .then(data => setValid(data === "valid"))
      .catch(() => setValid(false));
  }, []);

  const handleLogout = () => {
    fetch(`${baseURL}/logout`, { credentials: "include" })
      .then(() => setValid(false))
      .catch(() => setValid(false));
  };

  const handleDelete = () => {
    fetch(`${baseURL}/delete`, {
      method: "POST",
      credentials: "include"
    })
      .then(res => res.text())
      .then(data => {
        if (data === "deleted") {
          setValid(false);
          window.location.href = "/";
        } else {
          alert("Failed to delete account");
        }
      })
      .catch(() => alert("Server error"));
  };

  return (
    <div className="z-10 fixed navbar border-2 bg-blue-700 border-b-blue-700 flex justify-between">

      <a className="btn btn-ghost text-xl font-extrabold text-white" href="/">
        Roommate Finder
      </a>
      <div className="flex items-center gap-4">
        {!valid && (
          <>
            <a href="/login" className="btn btn-primary">Login</a>
          </>
        )}
        {valid && (
          <>
            <span className="text-white font-bold">Welcome back!</span>
            <button onClick={handleLogout} className="btn btn-error">Logout</button>
            <button onClick={() => setShowConfirm(true)} className="btn btn-warning">
              Delete Account
            </button>
          </>
        )}
      </div>

      {showConfirm && (
        <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50">
          <div className="bg-white p-6 rounded shadow-lg">
            <p className="mb-4">Are you sure you want to delete your account? This action cannot be undone.</p>
            <div className="flex gap-4">
              <button
                onClick={handleDelete}
                className="btn btn-error"
              >
                Yes, Delete
              </button>
              <button
                onClick={() => setShowConfirm(false)}
                className="btn btn-secondary"
              >
                Cancel
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}
