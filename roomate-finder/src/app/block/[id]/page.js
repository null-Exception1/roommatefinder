import Navbar from "@/components/navbar";
import RoomCard from "@/components/roomcard";
import Breadcumbs from "@/components/breadcrumbs";

export default async function Block({ params }) {
  const { id } = await params;
  const baseURL = process.env.NEXT_PUBLIC_API_URL;
  const res = await fetch(`${baseURL}/rooms?block=${id}`);
  const loaded_data = await res.json();

  return (
    <div className="min-h-screen bg-gray-50">
      <Navbar />
      <div className="flex justify-center pt-20 px-2 sm:px-4">
        <div className="flex flex-col my-4 sm:my-10 w-full max-w-7xl border-2 border-black rounded-3xl bg-white">
          <div>
            <Breadcumbs crumbs={["Home", "Blocks", `Block ${id}`]} links={["/", "/blocks", `/block/${id}`]} />
            <h1 className="pl-4 sm:pl-10 pt-5 text-xl sm:text-2xl font-bold">BLOCK {id}</h1>

            {loaded_data ? (
              <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4 p-4 sm:p-10 justify-items-center">
                {Object.entries(loaded_data).map(([roomID, roomData]) => (
                  <RoomCard
                    key={roomID}
                    RoomID={roomID}
                    People={roomData.People.map(p => ({
                      name: p.Name,
                      social: p.Social,
                      socialtype: p.Socialtype,
                    }))}
                  />
                ))}
              </div>
            ) : (
              <div className="flex justify-center items-center p-10">
                <span className="loading loading-spinner loading-lg"></span>
              </div>
            )}
          </div>
        </div>
      </div>
    </div>
  );
}
