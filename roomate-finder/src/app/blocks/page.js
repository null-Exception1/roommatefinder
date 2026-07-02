import Navbar from "@/components/navbar";
import Card from "@/components/card";
import Breadcumbs from "@/components/breadcrumbs";

export default function Blocks() {
  return (
    <div className="flex justify-center">
      <Navbar />

      <div className="flex flex-col m-20 flex-wrap w-11/12 border-2 border-black rounded-3xl">
        <div className="pt-10 pl-10">
          <fieldset className="fieldset max-w-xs text-black">
            <legend className="fieldset-legend text-black">Campus</legend>

            <select
              defaultValue="Main Campus"
              className="select bg-white text-black border-gray-300 focus:border-blue-500"
            >
              <option disabled={true}>Pick a campus</option>
              <option className="bg-white text-black">Main Campus</option>
              <option className="bg-white text-black">Bangalore</option>
              <option className="bg-white text-black">Jaipur</option>
            </select>
          </fieldset>
        </div>

        <Breadcumbs crumbs={["Home", "Blocks"]} links={["/", "/blocks"]} />

        <h1 className="pl-10 pt-2 text-2xl">Blocks</h1>

        <div className="grid grid-cols-3 p-5 *:gap-16 gap-16 ml-10">
          <Card BlockID="1" PartialCount="5" FullCount="5" />
          <Card BlockID="2" PartialCount="3" FullCount="7" />
          <Card BlockID="3" PartialCount="6" FullCount="4" />
          <Card BlockID="4" PartialCount="2" FullCount="8" />
          <Card BlockID="5" PartialCount="4" FullCount="6" />
          <Card BlockID="6" PartialCount="7" FullCount="3" />
          <Card BlockID="1" PartialCount="5" FullCount="5" />
          <Card BlockID="2" PartialCount="3" FullCount="7" />
          <Card BlockID="3" PartialCount="6" FullCount="4" />
          <Card BlockID="4" PartialCount="2" FullCount="8" />
          <Card BlockID="5" PartialCount="4" FullCount="6" />
          <Card BlockID="6" PartialCount="7" FullCount="3" />
        </div>

      </div>

    </div>);
}