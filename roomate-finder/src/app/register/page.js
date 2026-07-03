import Navbar from "@/components/navbar";
export default function Login() {
  return (
    <div>
      <Navbar />
      <div className="flex m-5 h-11/12 justify-center min-h-screen items-center">
        <fieldset className="fieldset bg-white-200 border-blue-300 rounded-box w-xs border p-4">

          <legend className="fieldset-legend text-black">Register</legend>

          <label className="label">Admission Number (will be hashed with firstname) * </label>
          <input type="email" className="input bg-white border-2 border-blue-700 outline-blue-700" placeholder="Admn no." />

          <label className="label">First name (Password) *</label>
          <input type="email" className="input bg-white border-2 border-blue-700 outline-blue-700" placeholder="First no." />

          <br />
          <select defaultValue="Social" className="select bg-white text-black">
            <option className="bg-white text-black">Discord</option>
            <option className="bg-white text-black">Instagram</option>
            <option className="bg-white text-black">WhatsApp</option>
          </select>
          <label className="label">Social</label>
          <input type="text" className="input bg-white border-2 border-blue-700 outline-blue-700" placeholder="User/Number" />
          <br />
          <label className="label">Block No.</label>
          <input type="text" className="input bg-white border-2 border-blue-700 outline-blue-700" placeholder="Block No." />

          <label className="label">Room No.</label>
          <input type="text" className="input bg-white border-2 border-blue-700 outline-blue-700" placeholder="Room No." />

          <button className="btn btn-info mt-4">Register</button>

        </fieldset>

      </div >
    </div>
  );
}