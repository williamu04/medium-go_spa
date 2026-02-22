'use client'

import { Box, Flex, Icon, Separator, Text } from "@chakra-ui/react"
import { LuBookmark, LuFileText, LuHouse, LuUserRound, LuUsersRound } from "react-icons/lu"
import { usePathname } from "next/navigation"
import Link from "next/link"

export const SideBar = () => {
  const pathname = usePathname()

  const menus = [
    { name: "Home", icon: LuHouse, href: "/" },
    { name: "Library", icon: LuBookmark, href: "/library" },
    { name: "Profile", icon: LuUserRound, href: "/profile" },
    { name: "Stories", icon: LuFileText, href: "/stories" },
    { name: "Following", icon: LuUsersRound, href: "/following" },
  ]

  return (
    <Box
      borderEndColor="gray.200"
      borderEndWidth="1px"
      width="40vh"
      height="100%"
    >
      <Flex direction="column" gap={2} mt={5}>
        {menus.slice(0, 4).map((menu) => {
          const isActive = pathname === menu.href

          return (
            <Link key={menu.name} href={menu.href}>
              <Flex
                align="center"
                gap={3}
                py={2}
                px={2}
                cursor="pointer"
                borderLeft={isActive ? "0.5px solid" : "transparent"}
                borderColor="black"
                fontWeight={isActive ? "bold" : "normal"}
                _hover={{ bg: "gray.50" }}
              >
                <Icon
                  as={menu.icon}
                  boxSize={5}
                  strokeWidth={isActive ? 2.5 : 1.5}
                />
                <Text>{menu.name}</Text>
              </Flex>
            </Link>
          )
        })}

        <Separator my={2} />

        {menus.slice(4).map((menu) => {
          const isActive = pathname === menu.href

          return (
            <Link key={menu.name} href={menu.href}>
              <Flex
                align="center"
                gap={3}
                py={2}
                px={2}
                cursor="pointer"
                borderLeft={isActive ? "0.5px solid" : "transparent"}
                borderColor="black"
                fontWeight={isActive ? "bold" : "normal"}
                _hover={{ bg: "gray.50" }}
              >
                <Icon
                  as={menu.icon}
                  boxSize={5}
                  strokeWidth={isActive ? 2.5 : 1.5}
                />
                <Text>{menu.name}</Text>
              </Flex>
            </Link>
          )
        })}
      </Flex>
    </Box>
  )
}
