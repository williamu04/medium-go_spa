'use effect'

import { Avatar, Card, Flex, Heading, HStack, Icon, Image, Separator, Spacer, Text } from "@chakra-ui/react"
import { LuBookmark, LuCircleMinus, LuEllipsis, LuMessageCircle } from "react-icons/lu"

interface Props {
  authorAvatar: string
  authorName: string
  title: string
  description: string
  thumbnail: string
  altThumbnail: string
  date: string
  commentCount: string
}

export const ArticleCard = ({
  authorAvatar, authorName, title, description, thumbnail, altThumbnail, date, commentCount
}: Props) => {
  return (
    <HStack>
      <Flex direction={"column"} flex={"1"}>

        <HStack>
          <Avatar.Root>
            <Avatar.Image src={authorAvatar} />
            <Avatar.Fallback name={authorName} />
          </Avatar.Root>
          <Text>{authorName}</Text>
        </HStack>

        <Flex direction={"column"} flex={"1"}>
          <Heading>{title}</Heading>
          <Text>{description}</Text>
        </Flex>

        <Flex>
          <Flex>
            <Text>{date}</Text>
            <Flex>
              <Icon as={LuMessageCircle} boxSize={"4"} />
              <Text>{commentCount}</Text>
            </Flex>
          </Flex>

          <Spacer />

          <Flex>
            <Icon as={LuCircleMinus} boxSize={"4"} />
            <Icon as={LuBookmark} boxSize={"4"} />
            <Icon as={LuEllipsis} boxSize={"4"} />
          </Flex>
        </Flex>

      </Flex>

      <Image src={thumbnail} alt={altThumbnail} height={"100px"} aspectRatio={3 / 2} />

    </HStack>
  )
}
