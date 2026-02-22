'use client'

import { Avatar, Box, Flex, Heading, HStack, Icon, Input, InputGroup, Spacer } from "@chakra-ui/react"
import { LuBell, LuMenu, LuSearch, LuSquarePen } from "react-icons/lu";

export const TopBar = () => {
  return (

    <Box w={"100%"} paddingX={"4"} paddingY={"1"} borderBottomWidth={"1px"} borderBottomColor={"gray"}>
      <Flex h="60px">
        <HStack gap={"5"}>
          <Box>
            <HStack>
              <Icon as={LuMenu} boxSize={"5"} />
              <Heading>MEDIUM</Heading>
              <InputGroup startElement={<LuSearch />}>
                <Input placeholder="Cari sesuatu" />
              </InputGroup>
            </HStack>
          </Box>
        </HStack>
        <Spacer />
        <HStack gap={"5"}>
          <Icon as={LuSquarePen} boxSize={"5"} />
          <Icon as={LuBell} boxSize={"5"} />
          <Avatar.Root size={"xs"}>
            <Avatar.Fallback name="Dunhill William" />
            <Avatar.Image src={"https://i.pravatar.cc/"} />
          </Avatar.Root>
        </HStack>
      </Flex>
    </Box>
  );
}


