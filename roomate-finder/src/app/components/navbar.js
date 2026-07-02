
export default function Navbar() {
  return (
    <div className="z-10 fixed navbar border-2 bg-blue-700 border-b-blue-700 flex justify-between">
      <a className="btn btn-ghost text-xl font-extrabold text-white" href="/">Roommate Finder</a>
      <div className="flex items-center gap-4">
        <a href="/login" className="btn btn-primary">Login</a>
      </div>
    </div >
  );
}