import Navbar from "./components/navbar";

export default function Home() {
  return (
    <div>
      <Navbar />

      <div className="flex h-screen max-h-screen justify-center min-h-screen items-center gap-4 bg-blue-100">

        <div className="card bg-white shadow h-1/2 w-1/3">
          <div className="card-body flex flex-col justify-evenly text-center items-center">
            <h1 className="text-5xl">Add yourself</h1>
            <h5 className="text-center text-2xl">alright bro what do you wanna do add yourself?</h5>
            <div className="justify-center card-actions">
              <a className="btn btn-primary" href="/register">Register</a>
            </div>
          </div>
        </div>

        <div className="card bg-white shadow h-1/2 w-1/3">
          <div className="card-body flex flex-col justify-evenly text-center items-center">
            <h1 className="text-5xl">Search</h1>
            <h5 className="text-center text-2xl">browse the blocks broski</h5>
            <div className="justify-center card-actions">
              <a className="btn btn-primary" href="/blocks">Browse</a>
            </div>
          </div>
        </div>

      </div>

    </div >
  );
}
