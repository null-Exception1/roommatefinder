/* eslint-disable @next/next/no-img-element */

export default function RoomCard({ RoomID, People }) {
  //const { people } = People;
  const badgeColors = {
    Instagram: "badge badge-dash badge-secondary",   // pink vibe
    Discord: "badge badge-dash badge-primary",   // discord purple
    Whatsapp: "badge badge-dash badge-accent",  // whatsapp green
    Twitter: "badge badge-dash badge-accent",
  };
  const icons = {
    Instagram: (
      <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"
        fill="currentColor" className="w-4 h-4 mr-1">
        <path d="M7 2C4.243 2 2 4.243 2 7v10c0 2.757 2.243 5 5 5h10c2.757 0 5-2.243 5-5V7c0-2.757-2.243-5-5-5H7zm10 2a3 3 0 013 3v10a3 3 0 01-3 3H7a3 3 0 01-3-3V7a3 3 0 013-3h10zm-5 3a5 5 0 100 10 5 5 0 000-10zm0 2a3 3 0 110 6 3 3 0 010-6zm4.5-2.75a1.25 1.25 0 11-2.5 0 1.25 1.25 0 012.5 0z" />
      </svg>
    ),
    Discord: (
      <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 245 240"
        fill="currentColor" className="w-4 h-4 mr-1">
        <path d="M104.4 104.9c-5.7 0-10.2 5-10.2 11.1s4.6 11.1 10.2 11.1c5.7 0 10.2-5 10.2-11.1s-4.6-11.1-10.2-11.1zm36.2 0c-5.7 0-10.2 5-10.2 11.1s4.6 11.1 10.2 11.1c5.7 0 10.2-5 10.2-11.1s-4.6-11.1-10.2-11.1z" />
        <path d="M189.5 20h-134C24.3 20 20 24.3 20 30v180c0 5.7 4.3 10 10 10h114l-5.4-18.8 13 12.1 12.3 11.4 21.6 19.3V30c0-5.7-4.3-10-10-10z" />
      </svg>
    ),
    Whatsapp: (
      <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32"
        fill="currentColor" className="w-4 h-4 mr-1">
        <path d="M16 3C9.4 3 4 8.4 4 15c0 2.6.8 5 2.2 7L4 29l7-2.2c2 1.4 4.4 2.2 7 2.2 6.6 0 12-5.4 12-12S22.6 3 16 3zm0 22c-2.2 0-4.3-.7-6-2l-.4-.3-4.2 1.3 1.4-4.1-.3-.4c-1.3-1.7-2-3.8-2-6 0-5.5 4.5-10 10-10s10 4.5 10 10-4.5 10-10 10z" />
        <path d="M21.6 18.4c-.3-.2-1.7-.9-2-1-.3-.1-.5-.2-.7.2-.2.3-.8 1-.9 1.2-.2.2-.3.3-.6.1-.3-.2-1.3-.5-2.5-1.6-.9-.8-1.6-1.8-1.8-2.1-.2-.3 0-.5.1-.7.1-.1.3-.3.4-.5.1-.2.2-.3.3-.5.1-.2.1-.4 0-.6-.1-.2-.7-1.7-1-2.3-.3-.6-.5-.5-.7-.5h-.6c-.2 0-.5.1-.7.3-.2.2-.9.9-.9 2.2s.9 2.6 1 2.8c.1.2 1.8 2.7 4.4 3.8.6.3 1.1.5 1.5.6.6.2 1.1.2 1.5.1.5-.1 1.7-.7 2-1.3.3-.6.3-1.1.2-1.3-.1-.2-.3-.3-.6-.5z" />
      </svg>
    ),
    Twitter: (
      <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32"
        fill="currentColor" className="w-4 h-4 mr-1">
        <path d="M16 3C9.4 3 4 8.4 4 15c0 2.6.8 5 2.2 7L4 29l7-2.2c2 1.4 4.4 2.2 7 2.2 6.6 0 12-5.4 12-12S22.6 3 16 3zm0 22c-2.2 0-4.3-.7-6-2l-.4-.3-4.2 1.3 1.4-4.1-.3-.4c-1.3-1.7-2-3.8-2-6 0-5.5 4.5-10 10-10s10 4.5 10 10-4.5 10-10 10z" />
        <path d="M21.6 18.4c-.3-.2-1.7-.9-2-1-.3-.1-.5-.2-.7.2-.2.3-.8 1-.9 1.2-.2.2-.3.3-.6.1-.3-.2-1.3-.5-2.5-1.6-.9-.8-1.6-1.8-1.8-2.1-.2-.3 0-.5.1-.7.1-.1.3-.3.4-.5.1-.2.2-.3.3-.5.1-.2.1-.4 0-.6-.1-.2-.7-1.7-1-2.3-.3-.6-.5-.5-.7-.5h-.6c-.2 0-.5.1-.7.3-.2.2-.9.9-.9 2.2s.9 2.6 1 2.8c.1.2 1.8 2.7 4.4 3.8.6.3 1.1.5 1.5.6.6.2 1.1.2 1.5.1.5-.1 1.7-.7 2-1.3.3-.6.3-1.1.2-1.3-.1-.2-.3-.3-.6-.5z" />
      </svg>
    )
  };

  return (
    <div className="card w-90 bg-white shadow-md hover:shadow-xl m-6 border border-black overflow-x-hidden">
      <div className="card-body flex flex-col items-center text-center">
        <h2 className="card-title mb-4">ROOM {RoomID}</h2>
        <div className="grid grid-rows-4 gap-3 text-sm w-full">
          {People.map((person, index) => (
            <div key={index} className="flex justify-center gap-3 flex-wrap">
              <span className="badge badge-neutral text-xs truncate">
                {person.name}
              </span>
              <span className={`${badgeColors[person.socialtype]} text-xs truncate flex flex-row`}>
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