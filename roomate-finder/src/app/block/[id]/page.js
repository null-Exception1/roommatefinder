
import Navbar from "@/components/navbar";
import RoomCard from "@/components/roomcard";
import Breadcumbs from "@/components/breadcrumbs";

// This remains an async Server Component
export default async function Block({ params }) {
  const { id } = await params; // Await params directly on the server
  const baseURL = process.env.NEXT_PUBLIC_API_URL;
  const res = await fetch(`${baseURL}/rooms?block=${id}`);
  const loaded_data = await res.json();

  return (
    <div className="flex justify-center">
      <Navbar />
      <div className="flex flex-col m-20 flex-wrap w-11/12 border-2 border-black rounded-3xl">
        <div>
          <Breadcumbs crumbs={["Home", "Blocks", `Block ${id}`]} links={["/", "/blocks", `/block/${id}`]} />
          <h1 className="pl-10 pt-5 text-2xl">BLOCK {id}</h1>
          {loaded_data ? (
            <div className="grid grid-cols-3 p-10">
              {Object.entries(loaded_data).map(([roomID, roomData]) => (
                <RoomCard
                  key={roomID}
                  RoomID={roomID}
                  People={roomData.People.map(p => ({ name: p.Name, social: p.Social, socialtype: p.Socialtype }))}
                />
              ))}
            </div>
          ) : (<span className="loading loading-spinner loading-xl"></span>)}
        </div>
      </div>
    </div>
  );
}
