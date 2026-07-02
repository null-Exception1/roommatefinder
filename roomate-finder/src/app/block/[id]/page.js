import Navbar from "@/components/navbar";
import RoomCard from "@/components/roomcard";
import Breadcumbs from "@/components/breadcrumbs";
export default async function Block({ params }) {
  const { id } = await params;

  return (
    <div className="flex justify-center">
      <Navbar />
      <div className="flex flex-col m-20 flex-wrap w-11/12 border-2 border-black rounded-3xl">
        <div className="">
          <Breadcumbs crumbs={["Home", "Blocks", `Block ${id}`]} links={["/", "/blocks", `/block/${id}`]} />
          <h1 className="pl-10 pt-5 text-2xl">BLOCK {id}</h1>
          <div className="grid grid-cols-3 p-10">
            <RoomCard RoomID="1" People={[{ "name": "Shaurya", "social": "discordusername", "socialtype": "discord" }, { "name": "Shauryaaaa", "social": "9958813899", "socialtype": "whatsapp" }, { "name": "Shauryaaaa", "social": "9958813899", "socialtype": "whatsapp" }, { "name": "Shauryaaaa", "social": "discordusername", "socialtype": "discord" }]} />
            <RoomCard RoomID="1" People={[{ "name": "Shaurya", "social": "discordusername", "socialtype": "discord" }, { "name": "Shauryaaaa", "social": "9958813899", "socialtype": "whatsapp" }, { "name": "Shauryaaaa", "social": "9958813899", "socialtype": "whatsapp" }, { "name": "Shauryaaaa", "social": "discordusername", "socialtype": "discord" }]} />
            <RoomCard RoomID="1" People={[{ "name": "Shaurya", "social": "discordusername", "socialtype": "discord" }, { "name": "Shauryaaaa", "social": "9958813899", "socialtype": "whatsapp" }, { "name": "Shauryaaaa", "social": "9958813899", "socialtype": "whatsapp" }, { "name": "Shauryaaaa", "social": "discordusername", "socialtype": "discord" }]} />

          </div>
        </div>
      </div>
    </div>
  );
}
