/* eslint-disable @next/next/no-img-element */
export default function RoomCard({ RoomID, People }) {
  const badgeColors = {
    Instagram: "badge badge-secondary",            // Pink vibe 
    Discord: "badge badge-ghost border-gray-300",  // Removed blue background, now neutral light gray
    Whatsapp: "badge badge-accent",               // Whatsapp green 
    Twitter: "badge badge-info",                  // Twitter blue 
  };

  return (
    <div className="card w-full max-w-sm bg-white shadow-md hover:shadow-xl border border-black transition-shadow duration-200">
      <div className="card-body p-5 flex flex-col items-center text-center space-y-3">
        <h2 className="card-title text-lg sm:text-xl font-semibold">ROOM {RoomID}</h2>

        <div className="flex flex-col gap-3 w-full">
          {People.map((person, index) => (
            <div key={index} className="flex flex-row justify-between items-center gap-2 border-b border-gray-100 pb-2 last:border-0 last:pb-0" >
              <span className="badge badge-neutral text-xs truncate max-w-[110px] sm:max-w-[130px]">
                {person.name}
              </span>

              <span className={`${badgeColors[person.socialtype] || "badge"} text-xs truncate flex items-center max-w-[140px] sm:max-w-[160px] py-2`} >
                <img
                  src={`/icons/${person.socialtype.toLowerCase()}.svg`}
                  alt={person.socialtype}
                  className="w-3.5 h-3.5 mr-1 flex-shrink-0"
                />
                <span className="truncate">{person.social}</span>
              </span>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}
