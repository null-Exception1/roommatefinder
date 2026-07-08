/* eslint-disable @next/next/no-img-element */

export default function RoomCard({ RoomID, People }) {
  const badgeColors = {
    Instagram: "badge badge-secondary",   // pink vibe
    Discord: "badge badge-primary",       // discord purple
    Whatsapp: "badge badge-accent",       // whatsapp green
    Twitter: "badge badge-info",          // twitter blue
  };

  return (
    <div className="card w-full sm:w-72 bg-white shadow-md hover:shadow-xl m-4 border border-black">
      <div className="card-body flex flex-col items-center text-center space-y-3">
        <h2 className="card-title text-lg sm:text-xl">ROOM {RoomID}</h2>
        <div className="flex flex-col gap-2 w-full">
          {People.map((person, index) => (
            <div
              key={index}
              className="flex justify-center items-center gap-2 flex-wrap"
            >
              <span className="badge badge-neutral text-xs truncate max-w-[120px]">
                {person.name}
              </span>
              <span
                className={`${badgeColors[person.socialtype]} text-xs truncate flex items-center max-w-[140px]`}
              >
                <img
                  src={`/icons/${person.socialtype.toLowerCase()}.svg`}
                  alt={person.socialtype}
                  className="w-4 h-4 mr-1"
                />
                {person.social}
              </span>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}
