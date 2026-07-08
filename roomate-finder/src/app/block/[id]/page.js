import Navbar from "@/components/navbar";
import RoomCard from "@/components/roomcard";
import Breadcumbs from "@/components/breadcrumbs";

export default async function Block({ params }) {
  const { id } = await params;
  const baseURL = process.env.NEXT_PUBLIC_API_URL;
  const res = await fetch(`${baseURL}/rooms?block=${id}`);
  const loaded_data = await res.json();

  return (
    <div>
      {/* Navbar fixed at top */}
      <Navbar />

      {/* Add top padding so content clears navbar */}
      <div className="flex justify-center pt-20">
        <div className="flex flex-col m-4 sm:m-20 flex-wrap w-full sm:w-11/12 border-2 border-black rounded-3xl">
          <div>
            <Breadcumbs
              crumbs={["Home", "Blocks", `Block ${id}`]}
              links={["/", "/blocks", `/block/${id}`]}
            />
            <h1 className="pl-4 sm:pl-10 pt-5 text-xl sm:text-2xl">BLOCK {id}</h1>

            {loaded_data ? (
              <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4 p-4 sm:p-10">
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
                <span className="loading loading-spinner loading-xl"></span>
              </div>
            )}
          </div>
        </div>
      </div>
    </div>
  );
}
