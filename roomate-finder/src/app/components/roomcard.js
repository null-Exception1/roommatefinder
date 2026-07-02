
export default function RoomCard({ RoomID, People }) {
  //const { people } = People;
  const badgeColors = {
    instagram: "badge badge-dash badge-secondary",   // pink vibe
    discord: "badge badge-dash badge-primary",   // discord purple
    whatsapp: "badge badge-dash badge-accent",   // whatsapp green
  };
  return (
    <div className="card w-80 bg-white shadow-md hover:shadow-xl m-6 border border-black overflow-x-hidden">
      <div className="card-body flex flex-col items-center text-center">
        <h2 className="card-title mb-4">ROOM {RoomID}</h2>
        <div className="grid grid-rows-4 gap-3 text-sm w-full">
          {People.map((person, index) => (
            <div key={index} className="flex justify-center gap-3 flex-wrap">
              <span className="badge badge-neutral text-xs truncate">
                {person.name}
              </span>
              <span className={`${badgeColors[person.socialtype]} text-xs truncate`}>
                {person.social}
              </span>
            </div>
          ))}
        </div>
      </div>
    </div>

  );
}