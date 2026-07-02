import Navbar from "@/components/navbar";
export default function Login() {
  return (
    <div>
      <Navbar />
      <div className="flex max-h-screen justify-center min-h-screen items-center">
        <fieldset className="fieldset bg-white-200 border-blue-300 rounded-box w-xs border p-4">
          <legend className="fieldset-legend text-black">Register</legend>

          <label className="label">Admission Number (will be hashed)</label>
          <input type="email" className="input bg-white border-2 border-blue-700 outline-blue-700" placeholder="Admn no." />

          <label className="label">Password</label>
          <input type="password" className="input bg-white border-2 border-blue-700 outline-blue-700" placeholder="Password" />

          <button className="btn btn-info mt-4">Login</button>

        </fieldset>

      </div >
    </div>
  );
}