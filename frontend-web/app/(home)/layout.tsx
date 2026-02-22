import { SideBar } from "@/components/layout/sidebar"
import { TopBar } from "@/components/layout/topbar"
import { Box, Flex } from "@chakra-ui/react"
import React from "react"

const HomeLayout = ({
  children,
}: {
  children: React.ReactNode;
}) => {
  return (
    <Flex h={"100vh"} direction={"column"}>
      <TopBar />
      <Flex flex={"1"} overflow={"hidden"}>
        <SideBar />
        <Box flex={"1"} overflowY={"auto"}>
          {children}
        </Box>
      </Flex>
    </Flex>
  );
}

export default HomeLayout;
